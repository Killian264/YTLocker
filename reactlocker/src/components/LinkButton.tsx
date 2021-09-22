import { Link } from "./Link";

export interface LinkButtonProps {
	className?: string;
	selected?: boolean;
	onClick: () => void;
}

export const LinkButton: React.FC<LinkButtonProps> = ({ className, selected = false, onClick, children }) => {
	let selectedCSS = selected ? "underline bg-primary-500" : "";

	return (
		<Link onClick={onClick} className={`${selectedCSS} text-lg font-semibold p-1 px-2 rounded`}>
			{children}
		</Link>
	);
};
