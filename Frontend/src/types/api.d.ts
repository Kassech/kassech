// src/types/api.d.ts
import { User } from "@/types/user";

// Type for the response from the API (both success and error format)
export interface ApiResponse<T = unknown> {
  status: "success" | "error";
  message: string;
  data: T | null;
}

// Type for the login credentials
export interface LoginCredentials {
  email_or_phone: string;
  password: string;
}

export interface LoginSuccessResponse {
    accessToken: string;
    user: User;
  }
export interface LoginErrorResponse {
    errorCode: string;
    errorMessage: string;
}

  export type MutationResponse<T = unknown> = T | undefined;


export interface Station {
  ID: number;
  LocationName: string;
  Latitude: number;
  Longitude: number;
}

// Type for the data used to create or update a station
export interface StationInput {
  LocationName: string;
  Latitude: number;
  Longitude: number;
}

// Type for the response when fetching a list of stations
export interface StationsResponse {
  stations: Station[];
}
