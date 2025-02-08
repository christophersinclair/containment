#include <linux/bpf.h>
#include <bpf/libbpf.h>
#include <unistd.h>

struct event {
    __u32 pid;
    char filename[512];
};

static int event_logger(void *ctx, void *data, size_t len)
{
    struct event *evt = (struct event *)data;
    printf("PID = %d and filenamd=%s\n", evt->pid, evt->filename);
    return 0;
}

int main()
{
    const char *file_path = "execvemon.o";
    const char *map_name = "ringbuf";
    const char *prog_name = "detect_execve";

    struct bpf_object *bpf_obj = bgf_object__open(file_path);
    if (!bpf_obj) {
        fprintf(stderr, "Error! Failed to load %s\n", file_path);
        return 1;
    }

    int err = bpf_object__load(bpf_obj);
    if (err) {
        fprintf(stderr, "Failed to load %s\n", file_path);
        return 1;
    }

    int rbFd = bpf_object__find_map_fd_by_name(bpf_obj, map_name);
    struct ring_buffer *ring_buf = ring_buffer__new(rbFd, event_logger, NULL, NULL);
    if (!ring_buf) {
        puts("Failed to create ring buffer");
        return 1;
    }
    
    struct bpf_program *prog = bpf_object__find_program_by_name(bpf_obj, prog_name);
    if (!prog) {
        fprintf(stderr, "Failed to find eBPF program\n");
        return 1;
    }

    bpf_program__attach(prog);
    
    while(1) {
        ring_buffer__consume(ring_buf);
        sleep(1);
    }

    return 0;
}