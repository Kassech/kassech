import { User } from './user';

export interface Driver extends User {
  drivingLicense?: File;
  nationalId?: File;
  insuranceDocument?: File;
  others?: File;
}
