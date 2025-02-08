import { Button } from '@/components/ui/button';
import { DriverDropDown } from './driverDropDown';
import { Card, CardHeader, CardTitle } from '@/components/ui/card';
import {VehicleDropDown} from './vehicleDropDown';

export default function DelegationForm() {

  const handleDriverSelect = (id: string) => {
    // setValue('driverID', id); // Use setValue from react-hook-form
  };

  const handleVehicleSelect = (id: string) => {
    // setValue('vehicleID', id); // Use setValue for vehicleID
  };

  return (
    <div className="flex flex-col items-center justify-center space-y-5">
      <Card className="py-8 px-4 w-full mx-2 flex flex-col items-center justify-center">
        <CardHeader className=" flex items-start justify-start">
          <CardTitle>Delegation</CardTitle>
        </CardHeader>
        <div className="md:pt-8 md:w-1/4 w-full pt-5">
          <DriverDropDown onDriverSelect={handleDriverSelect} />
        </div>

        <div className="md:pt-8 md:w-1/4 w-full pt-5">
          <VehicleDropDown onVehicleSelect={handleVehicleSelect} />
        </div>
        <Button type="submit" className="md:w-1/4 w-full rounded-lg mt-7">
          Submit
        </Button>
      </Card>
    </div>
  );
}
