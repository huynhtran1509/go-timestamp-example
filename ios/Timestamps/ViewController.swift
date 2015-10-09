//
//  ViewController.swift
//  Timestamps
//
//  Created by Ian Terrell on 10/9/15.
//  Copyright © 2015 WillowTree Apps. All rights reserved.
//

import UIKit
import Timestamp

struct TS: CustomStringConvertible {
    static var Formatter: NSDateFormatter = {
        let formatter = NSDateFormatter()
        formatter.dateStyle = NSDateFormatterStyle.LongStyle
        formatter.timeStyle = .MediumStyle
        return formatter
    }()
    
    let time: NSDate
    
    init(json: NSDictionary) {
        let s = json["seconds_since_epoch"] as! NSNumber
        time = NSDate(timeIntervalSince1970: s.doubleValue)
    }
    
    static func parse(jsonData: NSData) throws -> [TS] {
        let json = try NSJSONSerialization.JSONObjectWithData(jsonData, options: .AllowFragments) as! NSArray
        var timestamps = [TS]()
        for case let dict as NSDictionary in json {
            timestamps.append(TS(json: dict))
        }
        return timestamps
    }
    
    var description: String {
        return TS.Formatter.stringFromDate(time)
    }
}

class ViewController: UITableViewController {
    var timestamps = [TS]()
    
    override func viewDidLoad() {
        super.viewDidLoad()
        reloadTimestamps()
    }
    
    func reloadTimestamps() {
        var data: NSData?
        var error: NSError?
        Timestamp.GoTimestampAllJSON(&data, &error)
        guard error == nil else {
            print("Error getting timestamps: \(error)")
            return
        }
        
        do {
            timestamps = try TS.parse(data!)
        } catch {
            print("Error parsing timestamps: \(error)")
        }
        
        tableView.reloadData()
    }
    
    @IBAction func addTimestamp(sender: AnyObject) {
        var error: NSError?
        let now = NSDate().timeIntervalSince1970
        
        Timestamp.GoTimestampInsert(Int64(now), &error)
        guard error == nil else {
            print("Error inserting timestamp: \(error)")
            return
        }
        
        reloadTimestamps()
    }
    
    override func numberOfSectionsInTableView(tableView: UITableView) -> Int {
        return 1
    }
    
    override func tableView(tableView: UITableView, numberOfRowsInSection section: Int) -> Int {
        return timestamps.count
    }
    
    override func tableView(tableView: UITableView, cellForRowAtIndexPath indexPath: NSIndexPath) -> UITableViewCell {
        let cell = tableView.dequeueReusableCellWithIdentifier("TimestampCell")!
        cell.textLabel?.text = timestamps[indexPath.row].description
        return cell
    }

}

