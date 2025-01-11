export interface User {
    Id: number;
    Email: string;
    Password: string; // Hashed password
    FirstName: string;
    LastName: string;
    ProfilePicture?: string;
    ProfilePictureFile?: File;
    PhoneNumber: string;
    IsVerified: boolean;
    Role: number;
    IsOnline: boolean;
}
