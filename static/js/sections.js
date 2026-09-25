document.addEventListener("DOMContentLoaded", function () {

    const sections = document.querySelectorAll(
        ".side-section, .section, .experience, .formation"
    );

    sections.forEach(function (section) {

        section.addEventListener("mouseenter", function () {
            section.classList.add("active");
        });

        section.addEventListener("mouseleave", function () {
            section.classList.remove("active");
        });

    });

});