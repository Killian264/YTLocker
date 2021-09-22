import { Link } from "./Link";

export const Footer: React.FC<{}> = () => {
	return (
		<div className="max-w-7xl mx-auto w-full">
			<div className="bg-primary-700 opacity-50 pt-4 pb-4 mx-4 rounded-t-md">
				<div className="flex justify-between px-6 max-w-7xl mx-auto">
					<div className="flex gap-8">
						<Link
							className="text-lg font-semibold"
							href="https://www.youtube.com/watch?v=dQw4w9WgXcQ"
							target="_blank"
						>
							Terms
						</Link>
						<Link
							className="text-lg font-semibold"
							href="https://github.com/Killian264/YTLocker/issues"
							target="_blank"
						>
							Issues
						</Link>
						<Link
							className="text-lg font-semibold"
							href="https://github.com/Killian264/YTLocker"
							target="_blank"
						>
							Github
						</Link>
					</div>
					<Link
						className="text-lg font-semibold"
						href="https://killiandebacker.com/"
						target="_blank"
					>
						Portfolio
					</Link>
				</div>
			</div>
		</div>
	);
};
