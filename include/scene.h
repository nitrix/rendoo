#ifndef RENDOO_SCENE_H
#define RENDOO_SCENE_H

#include <stdio.h>
#include <stdlib.h>

#include "toolkit.h"

struct scene {
    struct list *objects;
};

struct scene *scene_create(void);
void scene_destroy(struct scene *scene);
void scene_clear(struct scene *scene);
int scene_load(struct scene *scene, const char *filepath);

#endif
