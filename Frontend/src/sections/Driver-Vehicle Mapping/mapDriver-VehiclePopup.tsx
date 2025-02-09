import { useForm } from 'react-hook-form';
import { Button } from '@/components/ui/button';
import {
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogContent,
  DialogTitle,
} from '@/components/ui/dialog';
import { Label } from '@/components/ui/label';
import VehicleDropDown from './vehicleDropDown';
import { DriverDropDown } from './driverDropDown';

export default function DialogDemo() {
  const { setValue, handleSubmit } = useForm(); // Initialize useForm

  const handleVehicleSelect = (id: string, name: string) => {
    setValue('Vehicle', { id, name });
  };

   const handleDriverSelect = (id: string, name: string) => {
     setValue('Driver', { id, name });
   };

  const onSubmit = (data: any) => {
    console.log('Form submitted:', data);
  };

  return (
    <>
      <div>
        <DialogHeader>
          <DialogTitle>Assign Driver</DialogTitle>
          <DialogDescription>
            Select a registered driver, assign a vehicle, and save or cancel the
            changes.
          </DialogDescription>
        </DialogHeader>
        <DialogContent>
          <div className="grid gap-4 py-1">
            <form onSubmit={handleSubmit(onSubmit)}>
              <div className="grid w-full items-center gap-4">
                <div className="flex flex-col space-y-2 w-full">
                  <Label htmlFor="name">Registered Driver</Label>
                  <DriverDropDown onDriverSelect={handleDriverSelect} />
                </div>
                <div className="flex flex-col space-y-2 pb-2">
                  <Label htmlFor="framework">Registered Vehicle</Label>
                  <VehicleDropDown onVehicleSelect={handleVehicleSelect} />
                  {/* Capitalized */}
                </div>
              </div>
            </form>
          </div>
        </DialogContent>
        <DialogFooter>
          <Button type="submit">Save changes</Button>{' '}
          {/* Submit button inside form */}
        </DialogFooter>
      </div>
    </>
  );
}
