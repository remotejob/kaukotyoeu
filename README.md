# Kauko Työ 
kaukotyoeu

## Project demonstrate real application installed on Google Container Engine

* Working link [kaukotyo.eu](http://kaukotyo.eu/)

* Main difference from previous project [Link](https://github.com/remotejob/clusters_export/tree/master/docker-kaukotyo)
in use golang web server instead of Nginx.

## First component it's server.go file
* static contents as well as templates (assets,templates dirs)  incorporated in docker container
by command COPY in Dockerfile
* COPY assets/ /assets/
* COPY templates /templates/
* We are need only one line in code to serve all assets contents r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fs))
* to make some visual effects used small JavaScript fragment in templates files (home_page.html, layout.html)
