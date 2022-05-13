function dropdown() {
  document.getElementById(this.name).classList.toggle("show");
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

drops = document.getElementsByClassName("dropdownButton");
for (i = 0; i < drops.length; i++) {
	drops[i].addEventListener('click', dropdown, false);	
}
