const hamburger = document.querySelector(".hamburger");
const navMenu = document.querySelector(".nav-menu");
const form = document.querySelector(".form-container");

hamburger.addEventListener("click", mobileMenu);

function mobileMenu() {
  hamburger.classList.toggle("active");
  navMenu.classList.toggle("active");
  form.classList.toggle("margin-extra");
}



$(document).ready(function() {
  $('#datepicker').datepicker({
    startDate: new Date(),
    endDate: '+10d',
    multidate: true,
    format: "dd-mm-yyyy",
    language: 'en'
  }).on('changeDate', function(e) {
    // `e` here contains the extra attributes
    $(this).find('.input-group-addon .count').text(' ' + e.dates.length);
  });
});
