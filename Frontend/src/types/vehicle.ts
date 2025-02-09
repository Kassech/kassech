export type Vehicle = {
  id: number;
  vin: string;
  licensePlateNumber: string;
  make: string;
  year: string;
  color: string;
  carType: string;
  carPicture: File | null;
  bollo: File | null;
  insurance: File | null;
  libre: File | null;
  ownerId: number;
  createdAt: string;
  updatedAt: string;
};
