#ifndef RENDOO_RASTERIZER_H
#define RENDOO_RASTERIZER_H

#include "toolkit.h"
#include "scene.h"

struct rasterizer {
    int foo;
};

struct rasterizer *rasterizer_create(void);
void rasterizer_destroy(struct rasterizer *rasterizer);

struct image *rasterizer_render(struct rasterizer *rasterizer, struct scene *scene);

#endif
