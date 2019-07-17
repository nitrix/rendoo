#include "scene.h"

struct scene *scene_create(void) {
    struct scene *scene = malloc(sizeof *scene);

    if (scene) {
        scene->objects = list_create();
    }

    return scene;
}

void scene_destroy(struct scene *scene) {
    scene_clear(scene);

    list_destroy(scene->objects);
    free(scene);
}

void scene_clear(struct scene *scene) {
    list_clear(scene->objects);
}

int scene_load(struct scene *scene, const char *filepath) {
    // TODO
    return -1;
}
