import React from 'react';

interface PathSelectionSectionProps {
  paths: { path_id: number }[];
  onPathSelect: (pathId: number) => void;
}

const PathSelectionSection: React.FC<PathSelectionSectionProps> = ({
  paths,
  onPathSelect,
}) => {
  return (
    <div className="p-4 bg-gray-50 rounded-lg shadow-sm">
      <h2 className="text-lg font-semibold mb-4">Select a Path</h2>
      <select
        className="p-2 border rounded w-full"
        onChange={(e) => onPathSelect(Number(e.target.value))}
      >
        <option value="">Choose a Path</option>
        {paths.map((path) => (
          <option key={path.path_id} value={path.path_id}>
            Path {path.path_id}
          </option>
        ))}
      </select>
    </div>
  );
};

export default PathSelectionSection;
