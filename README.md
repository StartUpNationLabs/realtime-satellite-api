<a id="readme-top"></a>

[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![MIT License][license-shield]][license-url]
<br />
<div align="center">
  <a href="https://github.com/StartUpNationLabs/react-flight-tracker-satellite">
    <img src="https://upload.wikimedia.org/wikipedia/commons/c/cf/Transiting_Exoplanet_Survey_Satellite_artist_concept_%28black_background%29.png" >
  </a>

<h3 align="center">Satellite Information API</h3>

  <p align="center">
    An API to calculate and retrieve information about satellites orbiting the Earth.
    <br />
    <a href="https://github.com/StartUpNationLabs/react-flight-tracker-satellite"><strong>Explore the docs »</strong></a>
    <br />
    <br />
    <a href="https://github.com/StartUpNationLabs/react-flight-tracker-satellite">View Demo</a>
    ·
    <a href="https://github.com/StartUpNationLabs/react-flight-tracker-satellite/issues">Report Bug</a>
    ·
    <a href="https://github.com/StartUpNationLabs/react-flight-tracker-satellite/issues">Request Feature</a>
  </p>
</div>

<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#api-endpoints">API Endpoints</a></li>
    <li><a href="#roadmap">Roadmap</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
    <li><a href="#acknowledgments">Acknowledgments</a></li>
  </ol>
</details>

<!-- ABOUT THE PROJECT -->

## About The Project

This project provides a comprehensive API to calculate and retrieve information about satellites orbiting the Earth. It
includes endpoints to get detailed satellite information, groups, minimal data, and positions.

Here's why:

* Accurate and up-to-date information on satellites.
* Simplified access to satellite data through RESTful endpoints.
* Easy integration into existing applications.

<p align="right">(<a href="#readme-top">back to top</a>)</p>

### Built With

* [Go](https://golang.org)

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- GETTING STARTED -->

## Getting Started

To get a local copy up and running, follow these simple steps.

### Prerequisites

* Go
  ```sh
  go install
  ```
* Buf (for protobuf)
  ```sh
  brew tap bufbuild/buf
  brew install buf
  ```
* Protoc (for protobuf)
  ```sh
  brew install protobuf
  ```
* Buf generate(To regenerate the protobuf files)
  ```sh
  buf generate
  ```


### Installation

1. Clone the repo
   ```sh
   git clone https://github.com/StartUpNationLabs/react-flight-tracker-satellite.git
   ```
2. Change to the project directory
   ```sh
   cd react-flight-tracker-satellite
   ```
3. Run the application
   ```sh
   go run cmd/main.go
   ```

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- USAGE EXAMPLES -->

## Usage

### Environment Variables

To run the application, you need to set the following environment variables:
  - `CELESTRACK_URL`: The URL to the Celestrak API.
  - `SPACETRACK_URL`: The URL to the Space-Track API.
  - `SPACETRACK_USERNAME`: The username for the Space-Track API.
  - `SPACETRACK_PASSWORD`: The password for the Space-Track API.
  - 'HOST': The host for the server(0.0.0.0 if in docker)

### Get Satellite Detail

```sh
GET /v1/satellite/detail/{id}
```

Response:

```json
{
  "ccsdsOmmVers": "1.0",
  "comment": "No comment",
  "creationDate": "2021-01-01T00:00:00Z",
  ...
}
```

_For more examples, please refer to
the [Documentation](https://github.com/StartUpNationLabs/react-flight-tracker-satellite)_

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- API ENDPOINTS -->

## API Endpoints

### Get Satellite Detail

* **URL**: `/v1/satellite/detail/{id}`
* **Method**: `GET`
* **Path Parameters**:
    * `id` (string): The ID of the satellite.
* **Success Response**:
    * **Code**: 200
    * **Content**: `v1SatelliteDetail`

### Get Satellite Groups

* **URL**: `/v1/satellite/groups`
* **Method**: `GET`
* **Success Response**:
    * **Code**: 200
    * **Content**: `v1GetSatelliteGroupsResponse`

### Get Minimal Satellites

* **URL**: `/v1/satellite/minimal`
* **Method**: `GET`
* **Query Parameters**:
    * `time` (string, optional): The time for which to get the satellite data.
    * `groups` (array, optional): The groups of satellites to retrieve.
* **Success Response**:
    * **Code**: 200
    * **Content**: `v1GetMinimalSatellitesResponse`

### Get Satellite Positions

* **URL**: `/v1/satellite/positions`
* **Method**: `GET`
* **Query Parameters**:
    * `time` (string, optional): The time for which to get the satellite positions.
    * `groups` (array, optional): The groups of satellites to retrieve positions for.
* **Success Response**:
    * **Code**: 200
    * **Content**: `v1GetSatellitePositionsResponse`

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- ROADMAP -->

## Roadmap

- [ ] Implement Satellite Path Calculation
- [ ] Add unit tests
- [ ] Add integration tests
- [ ] Improve error handling
- [ ] Enhance documentation

See the [open issues](https://github.com/StartUpNationLabs/react-flight-tracker-satellite/issues) for a full list of
proposed features (and known issues).

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- CONTRIBUTING -->

## Contributing

Contributions are what make the open source community such an amazing place to learn, inspire, and create. Any
contributions you make are **greatly appreciated**.

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also
simply open an issue with the tag "enhancement".
Don't forget to give the project a star! Thanks again!

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- LICENSE -->

## License

Distributed under the MIT License. See `LICENSE.txt` for more information.

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- CONTACT -->

## Contact

* Apoorva Srinivas Appadoo - [LinkedIn][linkedin-url]
* Axel Delille - [LinkedIn](https://www.linkedin.com/in/axel-delille/)

Project
Link: [https://github.com/StartUpNationLabs/react-flight-tracker-satellite](https://github.com/StartUpNationLabs/react-flight-tracker-satellite)

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- ACKNOWLEDGMENTS -->

## Acknowledgments

* [Space-Track](https://www.space-track.org/) for providing satellite data.
* [Celestrak](https://www.celestrak.com/) for providing satellite data.

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- MARKDOWN LINKS & IMAGES -->

[contributors-shield]: https://img.shields.io/github/contributors/StartUpNationLabs/react-flight-tracker-satellite.svg?style=for-the-badge

[contributors-url]: https://github.com/StartUpNationLabs/react-flight-tracker-satellite/graphs/contributors

[forks-shield]: https://img.shields.io/github/forks/StartUpNationLabs/react-flight-tracker-satellite.svg?style=for-the-badge

[forks-shield]: https://img.shields.io/github/forks/StartUpNationLabs/react-flight-tracker-satellite.svg?style=for-the-badge

[forks-url]: https://github.com/StartUpNationLabs/react-flight-tracker-satellite/network/members

[stars-shield]: https://img.shields.io/github/stars/StartUpNationLabs/react-flight-tracker-satellite?style=for-the-badge

[stars-url]: https://github.com/StartUpNationLabs/react-flight-tracker-satellite/stargazers

[issues-shield]: https://img.shields.io/github/issues/StartUpNationLabs/react-flight-tracker-satellite.svg?style=for-the-badge

[issues-url]: https://github.com/StartUpNationLabs/react-flight-tracker-satellite/issues

[license-shield]: https://img.shields.io/github/license/StartUpNationLabs/react-flight-tracker-satellite.svg?style=for-the-badge

[license-url]: https://github.com/StartUpNationLabs/react-flight-tracker-satellite/blob/master/LICENSE.txt

[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=for-the-badge&logo=linkedin&colorB=555

[linkedin-url]: https://www.linkedin.com/in/appadoo-apoorva-srinivas-481367207