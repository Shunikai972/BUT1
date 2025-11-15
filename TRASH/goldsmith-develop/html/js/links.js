document.addEventListener('DOMContentLoaded', () => {

  const trigger = document.querySelector('.distressed-link');
  const sphere = document.querySelector('.img-sphere');

  if (trigger && sphere) {
 /*    trigger.addEventListener('mouseenter', () => {
      sphere.src = "../images/arjowiggins.svg";
    }); */
    trigger.addEventListener('mouseleave', () => {
      sphere.src = "../images/1.svg";
    });
  }

});
