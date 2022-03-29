var slide_indexes = [];
var slideshows = document.querySelectorAll(".slideshowContainer");
for (let i = 0; i < slideshows.length; i++){
    slide_indexes[i] = [slideshows[i].id, 1];
    displaySlides(slide_indexes[i][1], i, slideshows[i].id);
}
function nextSlide(n, slideshow) {
    var i;
    for (i = 0; i<slide_indexes.length; i++){
	if (slide_indexes[i][0] == slideshow) {
	    break;
	}
    }
    displaySlides(slide_indexes[i][1] += n, i, slideshow);
}
function displaySlides(n, i, slideshow) {
    var slides = document.querySelectorAll('#' + slideshow + ' .showSlide');
    if (n > slides.length) { slide_indexes[i] = [slideshow, 1] }  
    if (n < 1) { slide_indexes[i] = [slideshow, slides.length] }  
    for (let x = 0; x < slides.length; x++) {  
        slides[x].style.display = "none";  
    } 
    slides[slide_indexes[i][1] - 1].style.display = "block";
}
