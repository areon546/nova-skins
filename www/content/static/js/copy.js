
var copiers = document.getElementsByClassName("copier");
var skins = document.getElementsByClassName("skin")

// functions 

var copy = function () {
  let textToCopy = this.getAttribute("csv")
  console.log(textToCopy)
  copyText(textToCopy)
};

function copyText(text) {
  navigator.clipboard.writeText(text)
}


// add to document

for (var i = 0; i < copiers.length; i++) {
  copiers[i].addEventListener('click', copy);
}
