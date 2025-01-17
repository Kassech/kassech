import { useForm } from 'react-hook-form';
import { Button } from '@/components/ui/button';
import {
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
} from '@/components/ui/dialog';
import { Label } from '@/components/ui/label';
import DriverDropDown from './driverDropDown';
import VehicleDropDown from './vehicleDropDown';

export default function DialogDemo() {
  const { setValue, handleSubmit } = useForm(); // Initialize useForm

  const handleOwnerSelect = (id: string, name: string) => {
    setValue('owner', { id, name });
  };

  const onSubmit = (data: any) => {
    console.log('Form submitted:', data);
  };

  return (
    <>
      <DialogHeader>
        <DialogTitle>Edit profile</DialogTitle>
        <DialogDescription>
          Make changes to your profile here. Click save when you're done.
        </DialogDescription>
      </DialogHeader>
      <div className="grid gap-4 py-4">
        <form onSubmit={handleSubmit(onSubmit)}>
          {' '}
          {/* Add submit handler */}
          <div className="grid w-full items-center gap-4">
            <div className="flex flex-col space-y-1.5">
              <Label htmlFor="name">Name</Label>
              <DriverDropDown onOwnerSelect={handleOwnerSelect} />{' '}
              {/* Capitalized */}
            </div>
            <div className="flex flex-col space-y-1.5">
              <Label htmlFor="framework">Framework</Label>
              <VehicleDropDown onOwnerSelect={handleOwnerSelect} />{' '}
              {/* Capitalized */}
            </div>
          </div>
          <DialogFooter>
            <Button type="submit">Save changes</Button>{' '}
            {/* Submit button inside form */}
          </DialogFooter>
        </form>
      </div>
    </>
  );
}
