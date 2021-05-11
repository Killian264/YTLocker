export interface SvgProps {
	size: number;
}

export const Plus: React.FC<SvgProps> = ({ size }) => {
	return (
		<svg width={size} height={size} fill="none" viewBox="0 0 24 24">
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

export const Arrow: React.FC<SvgProps> = ({ size }) => {
	return (
		<svg width={size} height={size} fill="none" viewBox="0 0 24 24">
			<path
				stroke="currentColor"
				strokeLinecap="round"
				strokeLinejoin="round"
				strokeWidth="1.5"
				d="M13.75 6.75L19.25 12L13.75 17.25"
			/>
			<path
				stroke="currentColor"
				strokeLinecap="round"
				strokeLinejoin="round"
				strokeWidth="1.5"
				d="M19 12H4.75"
			/>
		</svg>
	);
};
