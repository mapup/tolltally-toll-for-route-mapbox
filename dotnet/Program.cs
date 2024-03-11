using System;
using System.IO;
using System.Net;
using RestSharp;
namespace mapbox
{
    static class Constants
    {
        public const string MAPBOX_API_KEY = Environment.GetEnvironmentVariable("MAPBOX_API_KEY");
        public const string MAPBOX_API_URL = "https://api.mapbox.com/directions/v5/mapbox/driving";

        public const string TOLLGURU_API_KEY = Environment.GetEnvironmentVariable("TOLLGURU_API_KEY");
        public const string TOLLGURU_API_URL = "https://apis.tollguru.com/toll/v2";
        public const string POLYLINE_ENDPOINT = "complete-polyline-from-mapping-service";

        // Philadelphia, PA
        public const string source_longitude = "-96.7970";
        public const string source_latitude = "32.7767";
        // New York, NY
        public const string destination_longitude = "-74.0060";
        public const string destination_latitude = "40.7128";
    }

    public class Program
    {
        public static string get_Response(string source_latitude, string source_longitude, string destination_latitude, string destination_longitude)
        {
            string url = Constants.MAPBOX_API_URL + source_longitude + "," + source_latitude + ";" + destination_longitude + "," + destination_latitude + "?geometries=polyline&access_token=" + Constants.MAPBOX_API_KEY + "&overview=full";
            WebRequest request = WebRequest.Create(url);
            WebResponse response = request.GetResponse();
            String polyline;
            using (Stream dataStream = response.GetResponseStream())
            {
                StreamReader reader = new StreamReader(dataStream);
                string responseFromServer = reader.ReadToEnd();
                string[] output = responseFromServer.Split("\"geometry\":\"");
                string[] temp = output[1].Split("\"");
                polyline = temp[0];
            }
            response.Close();
            return (polyline);

        }
        public static string Post_Tollguru(string polyline)
        {
            var client = new RestClient(Constants.TOLLGURU_API_URL + "/" + Constants.POLYLINE_ENDPOINT);
            var request1 = new RestRequest(Method.POST);
            request1.AddHeader("content-type", "application/json");
            request1.AddHeader("x-api-key", Constants.TOLLGURU_API_KEY);
            request1.AddParameter("application/json", "{\"source\":\"mapbox\" , \"polyline\":\"" + polyline + "\" }", ParameterType.RequestBody);
            IRestResponse response1 = client.Execute(request1);
            var content = response1.Content;
            string[] result = content.Split("tag\":");
            string[] temp1 = result[1].Split(",");
            string cost = temp1[0];
            return cost;
        }
        public static void Main()
        {
            string polyline = get_Response(Constants.source_latitude, Constants.source_longitude, Constants.destination_latitude, Constants.destination_longitude);
            Console.WriteLine(Post_Tollguru(polyline));
        }
    }
}
