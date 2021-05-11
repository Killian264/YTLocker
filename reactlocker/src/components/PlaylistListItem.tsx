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
	return (
		<div
			className={`${className} font-semibold text-xl hover:bg-primary-600 rounded-md flex justify-between`}
		>
			<div className="flex p-1">
				<img
					src={playlist.thumbnail}
					alt="Logo"
					width="140"
					height="120"
					className="rounded-lg"
				/>
				<div className="pl-3 flex flex-col">
					<span>{playlist.title}</span>
					<Link
						className="text-accent text-lg ml-0.5"
						href={playlist.url}
						target="_blank"
					>
						Youtube Link
					</Link>
				</div>
			</div>
			<div className="mr-3 my-auto select-none">
				<Arrow size={24}></Arrow>
			</div>
		</div>
	);
};
