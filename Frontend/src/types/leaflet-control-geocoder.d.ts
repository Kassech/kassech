

declare module "leaflet-control-geocoder" {
  import * as L from "leaflet";

  namespace L.Control {
    interface GeocoderOptions {
      defaultMarkGeocode?: boolean;
    }

    function geocoder(options?: GeocoderOptions): L.Control;
  }

  export = L.Control;
}
