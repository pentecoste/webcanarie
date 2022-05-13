function dropdown(id) {
  document.getElementById(id).classList.toggle("show");
}

window.onclick = function(event) {
  if (!event.target.matches('.dropdownButton') && !event.target.matches('.dropdownButtonApartment')) {
    var dropdowns = document.getElementsByClassName("dropdownContent");
    var i;
    for (i = 0; i < dropdowns.length; i++) {
      var openDropdown = dropdowns[i];
      if (openDropdown.classList.contains('show')) {
        openDropdown.classList.remove('show');
      }
    }
  }
}
