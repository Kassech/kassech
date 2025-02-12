export interface User {
  id: number;
  email: string;
  password: string; // Hashed password
  first_name: string;
  last_name: string;
  Profile_picture?: string;
  ProfilePictureFile?: File;
  Phone_number: string;
  is_verified: boolean;
  roles: string[];
  IsOnline: boolean;
  permissions:string[];
}
