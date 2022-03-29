popup = {
 init: function(){
  $('a').click(function(){
    popup.open($(this));
  });

$(document).on('click', '.popup img', function(){
  return false;
}).on('click', '.popup', function(){
  popup.close();
})
},
open: function($a) {
$('.boxImg').addClass('pop');
$popup = $('<div class="popup" />').appendTo($('body'));
$fig = $a.clone().appendTo($('.popup'));
$close = $('<div class="close"><svg><use xlink:href="#close"></use></svg></div>').appendTo($fig);
src = $('img', $fig).attr('src');
setTimeout(function(){
  $('.popup').addClass('pop');
}, 10);
},
close: function(){
$('.boxImg, .popup').removeClass('pop');
setTimeout(function(){
  $('.popup').remove()
}, 100);
}
}

popup.init()
