package main

import (
	"math"
)

const (
	RAD_IN_DEG        float64 = math.Pi / 180
	EARTH_RADIUS_MEAN float64 = 6371008.8
	EARTH_RADIUS_A    float64 = 6378137.0 // https://en.wikipedia.org/wiki/Earth_radius#Equatorial_radius
	EARTH_RADIUS_B    float64 = 6356752.3 // https://en.wikipedia.org/wiki/Earth_radius#Polar_radius
	MERIDIAN_LENGTH   float64 = 20003930  // https://en.wikipedia.org/wiki/Meridian_(geography)
	EQUATOR_LENGTH    float64 = 40075160  // https://en.wikipedia.org/wiki/Equator
)

func deg2rad(degrees float64) float64 {
	return degrees * RAD_IN_DEG
}

func rad2deg(radians float64) float64 {
	return radians / RAD_IN_DEG
}

// func sq(x float64) float64 {
// 	return x * x
// }

//
// constexpr inline double meters_per_degree_y() {
//     return MERIDIAN_LENGTH / 180.0;
// }
//
// constexpr inline double meters_per_degree_x_rad(double latitude_rad) {
//     return EQUATOR_LENGTH / 360.0 * cos(latitude_rad);
// }
//
// constexpr inline double meters_per_degree_x_deg(double latitude_deg) {
//     return meters_per_degree_x_rad(deg2rad(latitude_deg));
// }
//
// earth_radius_rad(double lat) {
//     // https://en.wikipedia.org/wiki/Earth_radius#Location-dependent_radii
//     return std::sqrt(
//         (sq(sq(EARTH_RADIUS_A) * cos(lat)) + sq(sq(EARTH_RADIUS_B) * sin(lat)))
//         /
//         (sq(EARTH_RADIUS_A * cos(lat)) + sq(EARTH_RADIUS_B * sin(lat)))
//     );
// }
//
// constexpr inline double earth_radius_deg(double lat) {
//     return earth_radius_rad(deg2rad(lat));
// }
//
// EARTH_RADIUS_VIETNAM = earth_radius_deg(15);
//
// }

func meters_per_degree_lat() float64 {
	return MERIDIAN_LENGTH / 180
}

func meters_per_degree_lon(lat_deg float64) float64 {
    return EQUATOR_LENGTH / 360 * math.Cos(deg2rad(lat_deg));
}

func haversine(θ float64) float64 {
	return (1 - math.Cos(θ)) / 2
}

func distance_radians(lat1, lon1, lat2, lon2 float64) float64 {
	// returns distance between two points one a sphere of radius = 1
	Δlat := math.Abs(lat1 - lat2)
	Δlon := math.Abs(lon1 - lon2)
	if Δlat == 0 && Δlon == 0 {
		return 0
	}

	a := haversine(Δlat) + math.Cos(lat1)*math.Cos(lat2)*haversine(Δlon)
	return 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
}

func distance(lat1, lon1, lat2, lon2 float64) float64 {
	return EARTH_RADIUS_MEAN * distance_radians(
		deg2rad(lat1),
		deg2rad(lon1),
		deg2rad(lat2),
		deg2rad(lon2),
	)
}
