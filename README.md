
<h3 align="center">Nativefier Downloader</h3>


<p align="center">
  <a href="https://github.com/ghostx31/nativefier-downloader/stargazers"><img src="https://img.shields.io/github/stars/ghostx31/nativefier-downloader?colorA=363a4f&colorB=b7bdf8&style=for-the-badge"></a>
  <a href="https://github.com/ghostx31/nativefier-downloader/contributors"><img src="https://img.shields.io/github/contributors/ghostx31/nativefier-downloader?colorA=363a4f&colorB=a6da95&style=for-the-badge"></a> 
</p>

<img src="static/dist/assets/nativefier.webp">

Nativefier Downloader is a Golang and Tailwind CSS based project to convert your favourite websites into native-looking Electron apps!

Just enter the URL of your favourite website, select your OS and bam! You have an Electron app ready to use!

This project internally uses the [Nativefier NPM package](https://github.com/nativefier/nativefier) for building Electron apps.

### Contributing

If you wish to add a new feature, write your contributions on a new branch and open a PR against the dev branch.

### Building the Docker container 

- From the root of the repository run the command:
```bash
docker build --network=host -t nativefier:latest .
```

- To run the built docker container: 
```bash
docker run -p 1323:1323 nativefier:latest
```

- Now browse to `localhost:1323` to get to the page. 

- A docker image on dockerhub is also available. To use this image, run `docker run -p 1323:1323 spookyintheam/nativefier:latest`

### Helm Chart for Kubernetes

The repository also includes a helm chart with HPA for autoscaling. To build and deploy the helm chart, from the root of the project directory, run:

```bash
helm install nativefier nativefier-helm-chart/
```

Then get the external IP of the deployment's load balancer by running `kubectl get svc`. Open this IP address in the URL bar to access the project's homepage. 

### TODO

- [x] Support Electron apps for all three major OSes. 
- [x] Better frontend
- [x] Push image to dockerhub and also create a Kubernetes deployment Helm chart.
- [ ] Refine support for different versions of macOS.
- [ ] OS detection from browser (planned but not sure if I'll implement this).
