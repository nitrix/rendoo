#include "main.h"

int main(int argc, char *argv[]) {
    UNUSED(argc);
    UNUSED(argv);

	// Create the scene
	struct scene *scene = scene_create();
    DEFER(scene_destroy(scene));

    // Load an object
    scene_load(scene,"examples/african_head/object.ini");

	// Prepare the rasterizer
	struct rasterizer *rasterizer = rasterizer_create();
    DEFER(rasterizer_destroy(rasterizer));

    // Render the scene
    struct image *image = rasterizer_render(rasterizer, scene);
    DEFER(image_destroy(image));

    // Save the image to a file
    image_encode_png(image, "output.png");

    return EXIT_SUCCESS;
}
