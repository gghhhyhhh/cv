# CV Web — Hicham HEE

CV web professionnel développé avec **Go, HTML, CSS, JavaScript et Docker**.

## Fonctionnalités

* CV accessible depuis un navigateur
* Design responsive
* Organisation en plusieurs sections
* Photo et informations personnelles
* Présentation des compétences, expériences et projets
* Téléchargement du CV en **PDF A4**
* Déploiement avec **Docker**

## Technologies

* Go
* HTML
* CSS
* JavaScript
* Docker

## Lancement

```bash
cd ~/cursus/cv
docker build -t cv .
docker run -d --name cv -p 8080:8080 cv
```

Puis ouvrir :

```text
http://localhost:8080
```

## Arrêter le projet

```bash
docker stop cv
```

## Redémarrer

```bash
docker start cv
```

## Structure

```text
cv/
├── main.go
├── Dockerfile
├── docker-compose.yml
├── go.mod
├── downloads/
├── static/
│   ├── css/
│   ├── js/
│   └── photo.png
└── templates/
    ├── index.html
    └── downloads.html
```
