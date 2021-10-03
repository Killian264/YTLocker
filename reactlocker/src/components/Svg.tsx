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
		<div onClick={onClick} className={`${className} border-2 rounded-md cursor-pointer`}>
			{children}
		</div>
	);
};

export const ChevronDown: React.FC<{}> = () => {
	return (
		<svg
			xmlns="http://www.w3.org/2000/svg"
			width="16"
			height="16"
			fill="currentColor"
			className="bi bi-chevron-down"
			viewBox="0 0 16 16"
		>
			<path
				fillRule="evenodd"
				d="M1.646 4.646a.5.5 0 0 1 .708 0L8 10.293l5.646-5.647a.5.5 0 0 1 .708.708l-6 6a.5.5 0 0 1-.708 0l-6-6a.5.5 0 0 1 0-.708z"
			/>
		</svg>
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

export const Copy: React.FC<SvgProps> = ({ className, size }) => {
	return (
		<svg
			xmlns="http://www.w3.org/2000/svg"
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
				d="M8 16H6a2 2 0 01-2-2V6a2 2 0 012-2h8a2 2 0 012 2v2m-6 12h8a2 2 0 002-2v-8a2 2 0 00-2-2h-8a2 2 0 00-2 2v8a2 2 0 002 2z"
			/>
		</svg>
	);
};

export const Cancel: React.FC<SvgProps> = ({ className, size }) => {
	return (
		<svg
			xmlns="http://www.w3.org/2000/svg"
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
				d="M18.364 18.364A9 9 0 005.636 5.636m12.728 12.728A9 9 0 015.636 5.636m12.728 12.728L5.636 5.636"
			/>
		</svg>
	);
};

export const Pause: React.FC<SvgProps> = ({ className, size }) => {
	return (
		<svg
			xmlns="http://www.w3.org/2000/svg"
			width={size}
			height={size}
			fill="currentColor"
			className={className}
			viewBox="0 0 16 16"
		>
			<path d="M6 3.5a.5.5 0 0 1 .5.5v8a.5.5 0 0 1-1 0V4a.5.5 0 0 1 .5-.5zm4 0a.5.5 0 0 1 .5.5v8a.5.5 0 0 1-1 0V4a.5.5 0 0 1 .5-.5z" />
		</svg>
	);
};

export const Reload: React.FC<SvgProps> = ({ className, size }) => {
	return (
		<svg fill="#ffffff" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 30 30" width={size} height={size}>
			<path d="M 15 3 C 12.031398 3 9.3028202 4.0834384 7.2070312 5.875 A 1.0001 1.0001 0 1 0 8.5058594 7.3945312 C 10.25407 5.9000929 12.516602 5 15 5 C 20.19656 5 24.450989 8.9379267 24.951172 14 L 22 14 L 26 20 L 30 14 L 26.949219 14 C 26.437925 7.8516588 21.277839 3 15 3 z M 4 10 L 0 16 L 3.0507812 16 C 3.562075 22.148341 8.7221607 27 15 27 C 17.968602 27 20.69718 25.916562 22.792969 24.125 A 1.0001 1.0001 0 1 0 21.494141 22.605469 C 19.74593 24.099907 17.483398 25 15 25 C 9.80344 25 5.5490109 21.062074 5.0488281 16 L 8 16 L 4 10 z" />
		</svg>
	);
};
