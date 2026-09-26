document.addEventListener("DOMContentLoaded", function () {

    const projects = document.querySelectorAll(".project");

    projects.forEach(function (project) {

        project.addEventListener("click", function () {
            project.classList.toggle("selected");
        });

    });

});