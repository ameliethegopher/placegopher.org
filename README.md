# placegopher.org #

An image placeholder service using Gophers because ... well, why wouldn't you!

Note: As of 2017-08-18 the service is still in development, but will be launched soon.

## Using this Service ##

Anyone can use placegopher.org to show images of Gophers on their own designs or as placeholder images for whatever
reason. It's pretty easy to use, so just follow along.

There are currently three params you can tweak. Width is required, and height and color are optional:

* `w` (required) - the width of the image
* `h` (optional) - the height of the image (default: same as `w` to create a square image)
* `s` (optional) - saturation, choose 'c' (for 'color/colour') or 'g' (for 'grey/gray') (default: 'c')

Yes, I'm using 'query' params instead of path params. Why? Because it's totally fine and in this case much easier to
use for future expansion. Seriously. People should use them more often, esp. for uses like this.

Example URLs:

* https://placegopher.org/img?w=400&h=200        - colour picture of 400x200 pixels
* https://placegopher.org/img?w=512&c=gray - greyscale picture of 512 pixels square

Invalid URLs:

* https://placegopher.org/img?h=512        - missing width parameter `w`
* https://placegopher.org/img?w=sdf        - width is not a number
* https://placegopher.org/img?w=0          - width should be greater than 0
* https://placegopher.org/img?w=64&c=coolr - `coolr` is an invalid color/colour

## Redirecting to the Real Image ##

When an image is requested, we will do one of two things. Either:

* if an image already exists with the same params, we'll redirect to the image itself
* if an image of the specified criteria doesn't yet exist, it will be created, and then we'll redirect to the new image

Note: the images are located on a different sub-domain so that we have the option of using a CDN in the future. For
example your request to `/img` may be redirected to a url such as `https://cdn.placegopher.org/Pjx0NpF0507c`. This is
an opaque value and shouldn't be used directly. You should always use the `https://placegopher.org/img` endpoint in your
'img' tags or css.

## Choosing a Specific Image ##

Currently (2017-08-18) you are unable to choose a specific image when linking to a PlaceGopher image.

An image is chosen at random when we generate the first image of a particular size. e.g. we choose image number 10 when
generating an image of 400x200. After this point all requests for the same size (and same criteria) will use the same
image and won't change per request. (Note: we reserve the right to have to regenerate the images at any future point if
it'll make our lives easier, so just note you **might** not always get exactly the same image returned but you will get
the criteria fulfilled for what you request.)

## Future Features ##

These are all just possibilities:

* text on the image (e.g. `https://placegopher.org/img?h=512&t=Hello`)
* choose a specific image (e.g. `https://placegopher.org/img?h=512&i=10`)

## Author ##

Amelie the Gopher is a parody/joke/fake/random/misc account on both [Twitter](https://twitter.com/AmelieTheGopher) and
[GitHub](https://github.com/ameliethegopher/) run by someone mostly uninteresting. However, that same person feels like
contributing back to the Go community because of love, fondness, and sometimes they see something which the community
might like.

(Ends)
