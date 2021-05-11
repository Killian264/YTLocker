import React from "react";
import { Playlist } from "../shared/types";
import { Link } from "./Link";
import { Arrow } from "./Svg";

export interface PlaylistItemProps {
	className?: string;
	playlist: Playlist;
}

export const PlaylistListItem: React.FC<PlaylistItemProps> = ({
	playlist,
	className,
}) => {
	const css = `${className} hover:bg-primary-600 rounded-md flex justify-between`;

	const imageSize = "md:h-20 sm:h-16 h-16";

	const textSize = "sm:text-lg text-md";

	return (
		<div className={css}>
			<div className="flex p-1">
				<img
					src={playlist.thumbnail}
					alt="Logo"
					className={`rounded-lg object-cover ${imageSize}`}
				/>
				<div className="pl-3 flex flex-col">
					<span className={`${textSize} font-semibold`}>
						{playlist.title}
					</span>
					<Link
						className={`${textSize} text-accent ml-0.5`}
						href={playlist.url}
						target="_blank"
					>
						Youtube Link
					</Link>
				</div>
			</div>
			<div className="mr-2 my-auto select-none">
				<Arrow size={24}></Arrow>
			</div>
		</div>
	);
};