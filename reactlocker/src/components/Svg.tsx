export interface SvgProps {
	className?: string;
	size: number;
	strokeWidth?: number;
}

export interface SvgBoxProps {
	className?: string;
	onClick?: () => void;
}

export const SvgBox: React.FC<SvgBoxProps> = ({ className, onClick, children }) => {
	return (
		<div onClick={onClick} className={`${className} p-1 border-2 rounded-md cursor-pointer`}>
			{children}
		</div>
	);
};

export const Plus: React.FC<SvgProps> = ({ className, size }) => {
	return (
		<svg width={size} height={size} className={className} fill="none" viewBox="0 0 24 24">
			<path
				stroke="currentColor"
				strokeLinecap="round"
				strokeLinejoin="round"
				strokeWidth="1.9"
				d="M12 5.75V18.25"
			/>
			<path
				stroke="currentColor"
				strokeLinecap="round"
				strokeLinejoin="round"
				strokeWidth="1.9"
				d="M18.25 12L5.75 12"
			/>
		</svg>
	);
};

export const RightArrow: React.FC<SvgProps> = ({ size, strokeWidth = 1.5 }) => {
	return (
		<svg width={size} height={size} fill="none" viewBox="0 0 24 24">
			<path
				stroke="currentColor"
				strokeLinecap="round"
				strokeLinejoin="round"
				strokeWidth={strokeWidth}
				d="M13.75 6.75L19.25 12L13.75 17.25"
			/>
			<path
				stroke="currentColor"
				strokeLinecap="round"
				strokeLinejoin="round"
				strokeWidth={strokeWidth}
				d="M19 12H4.75"
			/>
		</svg>
	);
};

export const LeftArrow: React.FC<SvgProps> = ({ size, strokeWidth = 1.5 }) => {
	return (
		<svg width={size} height={size} fill="none" viewBox="0 0 24 24">
			<path
				stroke="currentColor"
				strokeLinecap="round"
				strokeLinejoin="round"
				strokeWidth={strokeWidth}
				d="M10.25 6.75L4.75 12L10.25 17.25"
			/>
			<path
				stroke="currentColor"
				strokeLinecap="round"
				strokeLinejoin="round"
				strokeWidth={strokeWidth}
				d="M19.25 12H5"
			/>
		</svg>
	);
};

export const Trash: React.FC<SvgProps> = ({ className, size }) => {
	return (
		<svg
			className={className}
			width={size}
			height={size}
			fill="none"
			viewBox={`0 0 24 24`}
			stroke="currentColor"
		>
			<path
				strokeLinecap="round"
				strokeLinejoin="round"
				strokeWidth={2}
				d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"
			/>
		</svg>
	);
};

export const Cog: React.FC<SvgProps> = ({ className, size }) => {
	return (
		<svg
			className={className}
			width={size}
			height={size}
			fill="none"
			viewBox="0 0 24 24"
			stroke="currentColor"
		>
			<path
				strokeLinecap="round"
				strokeLinejoin="round"
				strokeWidth={2}
				d="M10.325 4.317c.426-1.756 2.924-1.756 3.35 0a1.724 1.724 0 002.573 1.066c1.543-.94 3.31.826 2.37 2.37a1.724 1.724 0 001.065 2.572c1.756.426 1.756 2.924 0 3.35a1.724 1.724 0 00-1.066 2.573c.94 1.543-.826 3.31-2.37 2.37a1.724 1.724 0 00-2.572 1.065c-.426 1.756-2.924 1.756-3.35 0a1.724 1.724 0 00-2.573-1.066c-1.543.94-3.31-.826-2.37-2.37a1.724 1.724 0 00-1.065-2.572c-1.756-.426-1.756-2.924 0-3.35a1.724 1.724 0 001.066-2.573c-.94-1.543.826-3.31 2.37-2.37.996.608 2.296.07 2.572-1.065z"
			/>
			<path
				strokeLinecap="round"
				strokeLinejoin="round"
				strokeWidth={2}
				d="M15 12a3 3 0 11-6 0 3 3 0 016 0z"
			/>
		</svg>
	);
};

export const Checkmark: React.FC<SvgProps> = ({ className, size }) => {
	return (
		<svg
			className={className}
			fill="none"
			width={size}
			height={size}
			viewBox="0 0 24 24"
			stroke="currentColor"
		>
			<path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M5 13l4 4L19 7" />
		</svg>
	);
};

export const ExternalLink: React.FC<SvgProps> = ({ className, size }) => {
	return (
		<svg
			className={className}
			width={size}
			height={size}
			fill="none"
			viewBox="0 0 24 24"
			stroke="currentColor"
		>
			<path
				strokeLinecap="round"
				strokeLinejoin="round"
				strokeWidth={2}
				d="M10 6H6a2 2 0 00-2 2v10a2 2 0 002 2h10a2 2 0 002-2v-4M14 4h6m0 0v6m0-6L10 14"
			/>
		</svg>
	);
};
