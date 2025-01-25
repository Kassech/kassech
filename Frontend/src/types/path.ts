export interface Path {
    ID: number;
    CreatedAt: string;
    UpdatedAt: string;
    DeletedAt: string | null;
    route_id: number;
    route: {
      ID: number;
      StationA: {
        ID: number;
        LocationName: string;
        Latitude: number;
        Longitude: number;
      };
      StationB: {
        ID: number;
        LocationName: string;
        Latitude: number;
        Longitude: number;
      };
    };
    path_name: string;
    distance_km: number;
    estimated_time: string;
    is_active: boolean;
  }
