import React from "react";
import { Playlist } from "../shared/types";
import { ColorBadge } from "./ColorBadge";
import { RightArrow } from "./Svg";

export interface PlaylistItemProps {
	className?: string;
	playlist: Playlist;
	url: string;
	onClick: () => void;
}

export const PlaylistListItem: React.FC<PlaylistItemProps> = ({ playlist, className, url, onClick }) => {
	const css = `${className} hover:bg-primary-600 rounded-md flex justify-between cursor-pointer`;
	const imageSize = "md:h-20 sm:h-16 h-16";
	const textSize = "sm:text-md text-md";

	return (
		<div className={css} onClick={onClick}>
			<div className="flex p-1">
				<img
					src={playlist.thumbnailUrl}
					alt="Thumbnail"
					className={`rounded-lg object-cover ${imageSize}`}
				/>
				<div className="pl-3 flex flex-col">
					<span className={`${textSize} font-semibold`}>{playlist.title}</span>
					<span className="whitespace-nowrap">{playlist.description}</span>
					<div>
						<ColorBadge className="mt-2" color={playlist.color}></ColorBadge>
					</div>
				</div>
			</div>
			<div className="mr-2 my-auto select-none">
				<RightArrow size={24}></RightArrow>
			</div>
		</div>
	);
};
