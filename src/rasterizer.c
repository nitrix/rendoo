#include "rasterizer.h"

struct rasterizer *rasterizer_create(void) {
    struct rasterizer *rasterizer = malloc(sizeof *rasterizer);
    return rasterizer;
}

void rasterizer_destroy(struct rasterizer *rasterizer) {
    free(rasterizer);
}

struct image *rasterizer_render(struct rasterizer *rasterizer, struct scene *scene) {
    // TODO

    return NULL;
}
