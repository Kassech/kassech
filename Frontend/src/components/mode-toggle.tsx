import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from '@/components/ui/select';
import { useTheme } from './theme-provider';

export function ModeToggle() {
  const { setTheme } = useTheme();

  const handleThemeChange = (theme: string) => {
    setTheme(theme);
  };

  return (
    <Select defaultValue="system" onValueChange={handleThemeChange}>
      <SelectTrigger className="border-none">
        <SelectValue placeholder="Select Theme" />
      </SelectTrigger>
      <SelectContent>
        <SelectItem value="light">
          <span className="mr-2">🌞</span> Light
        </SelectItem>
        <SelectItem value="dark">
          <span className="mr-2">🌙</span> Dark
        </SelectItem>
        <SelectItem value="system">
          <span className="mr-2">💻</span> System
        </SelectItem>
      </SelectContent>
    </Select>
  );
}
