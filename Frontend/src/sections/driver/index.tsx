import React from 'react';
import DriverForm from './driverForm';
import DriverTable from './DriverTable';

const Driver: React.FC = () => {
  const switchTab = (tab: string) => {
    console.log(`Switching to tab: ${tab}`);
  };

  return <DriverForm switchTab={switchTab} />;
};

export default Driver;
