import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class SongListService {
  private apiUrl = 'http://localhost:8080/api';

  constructor(private http: HttpClient) {}

  getSongLists(): Observable<any[]> {
    return this.http.get<any[]>(`${this.apiUrl}/songlists`);
  }

  createSongList(songList: any): Observable<any> {
    return this.http.post<any>(`${this.apiUrl}/songlists`, songList);
  }

  deleteSongList(id: number): Observable<any> {
    return this.http.delete<any>(`${this.apiUrl}/songlists/${id}`);
  }
}
