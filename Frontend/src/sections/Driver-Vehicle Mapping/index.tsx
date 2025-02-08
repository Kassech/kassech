import React, { useState } from 'react';
import { DataTable } from './table/data-table';
import { columns } from './table/column';
import { useFetchUserData } from '../../services/userService';

import { Dialog, DialogContent, DialogTrigger } from '@/components/ui/dialog';

import AddPopup from './mapDriver-VehiclePopup';
import { Button } from '@/components/ui/button';

const Vehicle: React.FC = () => {
  const { data: users, isLoading, error } = useFetchUserData();
  console.log('🚀 ~ users:', users);

  if (isLoading) return <div>Loading...</div>;
  if (error) return <div>Error loading user data</div>;
  return (
    <div>
      <div className="p-6">
        <h1 className="text-xl font-semibold mb-4">Vehicle Delegation</h1>
        <Dialog>
          <DialogTrigger asChild className="">
            <Button>Assign Driver to Vehicle</Button>
          </DialogTrigger>
          <DialogContent className="sm:max-w-[425px]">
            <AddPopup />
          </DialogContent>
        </Dialog>
        <DataTable columns={columns} data={users.data ?? []} />
      </div>
    </div>
  );
};

export default Vehicle;
