# TollTally - Toll for Route (Mapbox)

This repository provides code examples to extend Mapbox mapping capabilities by adding toll information from [TollGuru](https://tollguru.com/) to the route information from Mapbox.

## What the product/repo is
A collection of SDK-like examples in multiple languages (JavaScript, Python, Go, PHP, Ruby, .NET) that demonstrate how to calculate toll costs for Mapbox-generated routes by integrating with the TollGuru API.

## Architecture
*   **Mapbox Integration**: Consumes the Mapbox Directions API to fetch route details and encoded polylines.
*   **TollGuru Wrapper**: Acts as a bridge to the TollGuru `complete-polyline-from-mapping-service` endpoint.
*   **Cross-Platform**: Provides native implementation examples in 6 different programming languages (JS, Python, Go, PHP, Ruby, .NET).
*   **Stateless Design**: No persistence layer is required; all operations are executed as real-time, on-demand API calls.
*   **Parameter Flexibility**: Supports passing complex vehicle configurations (axles, weight, height) to ensure accurate tolling for any vehicle type.
*   **Secure Communication**: Utilizes API Key/Token based authentication over HTTPS for all external requests.
*   **Output Mapping**: Demonstrates how to parse and map external API responses to useful application data.

Detailed information can be found in [docs/architecture.md](docs/architecture.md).

## Prerequisites
*   **Mapbox API Key**: Obtain from [Mapbox Dashboard](https://account.mapbox.com/).
*   **TollGuru API Key**: Obtain from [TollGuru Developers](https://tollguru.com/developers/docs/).
*   **Language Runtimes**: Node.js, Python, etc., depending on the example being used.

## Local Setup
1.  **Clone the repo**:
    ```bash
    git clone https://github.com/mapup/tolltally-toll-for-route-mapbox.git
    ```
2.  **Navigate to a language folder**:
    ```bash
    cd javascript # or python, go, etc.
    ```
3.  **Install dependencies**:
    *   JS: `npm install`
    *   Python: `pip install -r requirements.txt`
4.  **Configure Environment**:
    ```bash
    export MAPBOX_API_KEY='your_key'
    export TOLLGURU_API_KEY='your_key'
    ```

## How to run tests
Some language directories include a `Testing` or `test` folder with manual test scripts and sample data (e.g., Python and Go). For others, you can verify functionality by running the main entry point (e.g., `node index.js`) and observing the console output.

## How to deploy
These are demonstration scripts. For production, integrate the logic into your application services. Ensure secure management of API keys (e.g., using AWS Secrets Manager or environment variables).

## Where config lives
*   **Environment Variables**: `MAPBOX_API_KEY`, `TOLLGURU_API_KEY`.
*   **Parameters**: Vehicle and route parameters are defined within the scripts (e.g., `index.js`, `MapBox.py`).

## Known Limitations
*   Requires a valid polyline from Mapbox; errors in routing geometry will affect toll results.
*   Subject to rate limits of both Mapbox and TollGuru APIs.

## Maintenance and Operations
Refer to [docs/runbook.md](docs/runbook.md) for recovery steps, logging, and monitoring details.

---

## Key Features (TollGuru)
*   **Geographies**: Support for North America, Europe, Australia, Asia, and Latin America.
*   **Vehicle Types**: Support for Cars, Trucks, Taxis, Motorcycles, RVs, and more.
*   **Payment Options**: Rates for Tag, Cash, License Plate, Credit Card, and Prepaid.
*   **Time-based Tolls**: Support for "departure_time" to handle dynamic tolling and express lanes.
*   **Truck Specifics**: Support for height, weight, hazardous goods, and axle-based pricing.

For more examples, visit the [TollGuru API Parameter Examples](https://github.com/mapup/tollguru-api-parameter-examples/tree/main/request-bodies/02-Complete-Polyline-To-Toll).
