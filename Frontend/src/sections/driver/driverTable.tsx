import { useDrivers } from '../../services/driverService';

export default function DriverTable() {
    const { data: drivers } = useDrivers();

    return (
        <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
            {drivers?.map(driver => (
                <div key={driver.id} className="p-4 border rounded">
                    <h2 className="text-lg font-bold">{driver.name}</h2>
                    <p>{driver.email}</p>
                </div>
            ))}
        </div>
    );
}
