<?php

function getCord($address) {
    $key = getenv('MAPBOX_API_KEY');
    
    $url = 'https://api.mapbox.com/geocoding/v5/mapbox.places/' . urlencode($address) . '.json?access_token=' . $key . '';
    
    $ch = curl_init();
    
    curl_setopt($ch, CURLOPT_SSL_VERIFYHOST, false);
    curl_setopt($ch, CURLOPT_SSL_VERIFYPEER, false);
    
    curl_setopt($ch, CURLOPT_URL, $url);
    curl_setopt($ch, CURLOPT_RETURNTRANSFER, 1);
    
    $responseJson = curl_exec($ch);
    curl_close($ch);
    
    $response = json_decode($responseJson, true);
    
    $location = array(
        'x' => $response['features']['0']['center']['1'],
        'y' => $response['features']['0']['center']['0']
    );
    
    return $location;
}
?>