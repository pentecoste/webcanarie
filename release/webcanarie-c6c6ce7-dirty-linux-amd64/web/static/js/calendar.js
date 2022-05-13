document.getElementById("calendar0").style.display = "flex";

function prevMonth(index){
	if (index) {
		document.getElementById("calendar"+index).style.display = "none";
		document.getElementById("calendar"+(index-1)).style.display = "flex";
	}
}

function nextMonth(index){
	if (index < 11) {	
		document.getElementById("calendar"+index).style.display = "none";
		document.getElementById("calendar"+(index+1)).style.display = "flex";
	}
}
