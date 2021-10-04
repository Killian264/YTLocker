import React from "react";
import { Playlist } from "../shared/types";
import { Badge } from "./Badge";
import { ColorBadge } from "./ColorBadge";
import { RightArrow } from "./Svg";

export interface PlaylistItemProps {
	className?: string;
	playlist: Playlist;
	url: string;
	onClick: () => void;
}

export const PlaylistListItem: React.FC<PlaylistItemProps> = ({ playlist, className, url, onClick }) => {
	const css = `${className} hover:bg-primary-600 rounded-md flex justify-between cursor-pointer overflow-hidden`;
	const imageSize = "md:h-20 md:w-32 h-16 w-24";
	const textSize = "sm:text-md text-md";

	return (
		<div className={css} onClick={onClick}>
			<div className="flex p-1 overflow-hidden">
				<div className={`rounded-lg flex-shrink-0 bg-black ${imageSize}`} />
				<div className="pl-3 flex flex-col">
					<span className={`${textSize} font-semibold whitespace-nowrap`}>{playlist.title}</span>
					<span className="whitespace-nowrap whitespace-nowrap">{playlist.description}</span>
					<div className="flex">
						<Badge className="mt-2 mr-2" color="primary">
							YTLocker
						</Badge>
						<ColorBadge className="mt-2" color={playlist.color}></ColorBadge>
					</div>
				</div>
			</div>
			<div className="mr-2 ml-4 my-auto select-none">
				<RightArrow size={24}></RightArrow>
			</div>
		</div>
	);
};
