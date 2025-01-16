import { useTranslation } from 'react-i18next';
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from '@/components/ui/select';

export const LanguageSelector = () => {
  const { i18n } = useTranslation();

  const handleLanguageChange = (language: string) => {
    i18n.changeLanguage(language);
  };

  return (
    <Select defaultValue={i18n.language} onValueChange={handleLanguageChange}>
      <SelectTrigger className="border-none">
        <SelectValue placeholder="Select Language" />
      </SelectTrigger>
      <SelectContent>
        <SelectItem value="en">
          <span className="mr-2">🇬🇧</span> English
        </SelectItem>
        <SelectItem value="am">
          <span className="mr-2">🇪🇹</span> Amharic
        </SelectItem>
      </SelectContent>
    </Select>
  );
};
