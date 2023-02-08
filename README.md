
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
docker build --network=host -t nativefier-downloader:latest .
```

- To run the built docker container: 
```bash
docker run -p 1323:1323 nativefier-downloader:latest
```

- Now browse to `localhost:1323` to get to the page. 

### TODO

- [x] Support Electron apps for all three major OSes. 
- [x] Better frontend
- [ ] Refine support for different versions of macOS.
- [ ] OS detection from browser (planned but not sure if I'll implement this).
