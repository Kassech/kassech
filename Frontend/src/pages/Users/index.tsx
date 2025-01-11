import React from 'react';

import { useFetchUserData } from '../../services/userService'; // Replace with your actual service
import Header from '@/components/header';
import { DataTable } from './table/data-table';
import { columns } from './table/column';

const UserList: React.FC = () => {
  const { data: users, isLoading, error } = useFetchUserData();
  console.log('🚀 ~ users:', users);

  if (isLoading) return <div>Loading...</div>;
  if (error) return <div>Error loading user data</div>;

  const paths = [
    { name: 'Home', href: '/' },
    { name: 'Users', href: '/user' },
  ];

  return (
    <>
      <Header paths={paths} />
      <div className="p-6">
        <h1 className="text-xl font-semibold mb-4">User Management</h1>
        <DataTable columns={columns} data={users.users ?? []} />
      </div>
    </>
  );
};

export default UserList;
