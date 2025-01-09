import React from 'react';
import DriverForm from './driverForm';
import DriverTable from './driverTable';

const Driver: React.FC = () => {
  return (
    <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
      <DriverForm />
      <DriverTable />
    </div>
  );
};

export default Driver;
