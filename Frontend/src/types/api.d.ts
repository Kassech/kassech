// src/types/api.d.ts

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

// Type for the user data returned after a successful login
export interface User {
  id: number;
  firstName: string;
  lastName: string;
  email?: string | null;
  phoneNumber: string;
  isOnline: boolean;
  password: string;
  previousPassword1?: string | null;
  previousPassword2?: string | null;
  profilePicture?: string | null;
  isVerified: boolean;
  lastLoginDate?: Date | null;
  socketID?: string | null;
}

export interface LoginResponse {
  accessToken: string;
  user: User;
}

// Type for the mutation response (e.g., user data or undefined)
export type MutationResponse<T = unknown> = T | undefined;


export interface Station {
  id: number;
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