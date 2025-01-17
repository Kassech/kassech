export interface User {
  Id: number;
  email: string;
  password: string; // Hashed password
  first_name: string;
  last_name: string;
  ProfilePicture?: string;
  ProfilePictureFile?: File;
  PhoneNumber: string;
  IsVerified: boolean;
  Role: number;
  IsOnline: boolean;
}
