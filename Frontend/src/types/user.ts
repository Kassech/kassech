export interface User {
  ID: number;
  email: string;
  password: string; // Hashed password
  first_name: string;
  last_name: string;
  ProfilePicture?: string;
  ProfilePictureFile?: File;
  PhoneNumber: string;
  IsVerified: boolean;
  roles: string;
  IsOnline: boolean;
}
