import React, { useState } from 'react';
import AddPopup from './mapDriver-VehiclePopup';
import { Button } from '@/components/ui/button';
import { Dialog, DialogContent, DialogTrigger } from '@/components/ui/dialog';
import { DataTable } from './table/data-table';
import { columns } from './table/column';
import { useFetchUserData } from '../../services/userService'; 

const Vehicle: React.FC = () => {
 const { data: users, isLoading, error } = useFetchUserData();
 console.log('🚀 ~ users:', users);

 if (isLoading) return <div>Loading...</div>;
 if (error) return <div>Error loading user data</div>;
  return (
    <div>
      <div className="p-6">
        <h1 className="text-xl font-semibold mb-4">User Management</h1>
        <DataTable columns={columns} data={users.users ?? []} />
      </div>
      <Dialog>
        <DialogTrigger asChild className="fixed  left-1/2">
          <Button>Assign Driver to Vehicle</Button>
        </DialogTrigger>
        <DialogContent className="sm:max-w-[425px]">
          <AddPopup />
        </DialogContent>
      </Dialog>
    </div>
  );
};

export default Vehicle;
